# Billing System Overview

## Customer Assumptions

Customers want to:

1. **Predict costs**: Know if their balance will cover their workload across all products
2. **See per-product breakdown**: Understand charges per product by time period
3. **Control spending**: Adjust balance/resources when usage changes

## Product Decisions

### 1. Payments in USDC Only
The system uses a single payment token (USDC) hardcoded at deployment. Adding multiple currencies is complex and out of scope.

### 2. Anyone Can Pay for Any App
Anyone can deposit funds to pay for someone else's app via `depositFor(account, amount)`.

**Why segmented accounts matter**: Without restrictions, depositing to an account means the owner could:
- Withdraw your funds
- Use them to pay for other apps they own

**Solution - Segmented, restricted accounts**: Move apps to dedicated accounts that:
- Only pay for one specific app
- Cannot withdraw funds
- Cannot own other apps

This segmentation happens at the **account level** (not resource level) because tracking individual resources (apps, API keys, etc.) in a billing contract shared by multiple products is too gas-intensive/complex.

**Implementation options**:
1. **Smart contract account** - lacks withdrawal and multi-app ownership functions
2. **Protocol-level restrictions** - built-in controls on certain account types

**Example flow**:
1. Alice deploys app, bills to herself: `createApp(skuID, aliceAddress)`
2. App uses Compute + AI → both charge Alice's account
3. Admin transfers app to restricted account for community funding
4. All products now charge the restricted account
5. Community calls `depositFor(restrictedAccount, amount)` - funds locked to that app only

### 3. Single Account Model
One billing account pays for all products (Compute, AI, etc.). This is a product decision contingent on technical feasibility.

**Rationale**:
- Simplifies customer experience: one balance, one view of spending
- Enables cross-product cost analysis (see spending distribution)
- Single balance check shows position across all services
- Reduces cognitive load vs managing multiple accounts per product

**Technical enabler**: Hub-and-spoke architecture allows multiple product modules to charge the same account via `BillingCore.chargePeriod()`.

## System Architecture

### Hub-and-Spoke Design

**BillingCore (Hub):**
- Single source of truth for accounts and balances
- Each account has one `int96 balance` (positive = credit, negative = debt)
- Tracks all registered products (Compute = Product 1, AI = Product 2, etc.)
- Provides period calculation (calendar months from genesis)
- Accepts charges from product modules

**Product Modules (Spokes):**
- `ComputeBilling`: Time-based continuous accrual with period settlement
- `UsageBilling`: Metered discrete events with period settlement
- Each module calls `BillingCore.chargePeriod()` to deduct from account balance
- Multiple products can charge the same account

### Unified Settlement Pattern

Both ComputeBilling and UsageBilling use the same settlement pattern:
1. **Accumulate** charges during the period
2. **Settle** at period end (monthly)
3. **Charge** BillingCore for the period

This unified approach provides:
- Predictable monthly billing cycles
- Consistent behavior across all products
- Event-based period attribution for analytics
- Efficient batch settlement operations

### Key Data Structures

**Account:**
```solidity
struct Account {
    int96 balance;        // Can go negative (debt)
    uint96 totalSpent;    // Lifetime spending
    bool suspended;       // True when balance < 0
}
```

**Product:**
```solidity
struct Product {
    string name;
    address module;            // ComputeBilling, UsageBilling, etc.
    address revenueRecipient;  // Where revenue goes
    uint96 unclaimedRevenue;   // Revenue waiting to be claimed
    bool active;
}
```

**SKU (for Compute products):**
```solidity
struct SKU {
    uint96 runningRate;      // tokens/second when app is running
    uint96 stoppedRate;      // tokens/second when app is stopped
    uint16 vcpus;            // vCPUs required for resource tracking
    uint96 minimumDeposit;   // Prevents instant suspension
    string name;
    bool active;
}
```

**AccountState (for Compute billing):**
```solidity
struct AccountState {
    uint96 totalRunningRate;  // Sum of all running app rates
    uint96 totalStoppedRate;  // Sum of all stopped app rates
    uint40 lastUpdate;        // Last time charges were accrued
    uint96 accruedCharges;    // Charges waiting to be settled
    uint40 chargesStartTime;  // When current charges started accumulating
}
```

### How It Delivers on Customer Needs

#### 1. Predict Costs

**Current charges:**
```solidity
getOutstandingCharges(account) → uint96  // For ComputeBilling
getUnsettledUsage(account) → uint96      // For UsageBilling
```

**Period estimates:**
```solidity
getChargesForPeriod(account, period) → (amount, settled)
```

Customers can see accrued charges before they're settled and estimate how long their current balance will last.

#### 2. Per-Product Breakdown

**Period charges** (queries all product modules):
```solidity
getChargesForPeriod(account, period) → ProductCharges[] {
    productId, productName, amount, settled
}
```

Returns charges for each product for a specific period. Customers can see: "I spent 80% on Compute, 15% on AI, 5% on Storage last month."

**Calendar month periods**: Period 0 = deployment month, Period 1 = next month, etc. Natural accounting alignment.

#### 3. Control Spending

- **Prepaid model**: Deposit upfront, charges deduct automatically
- **Minimum deposits**: Each SKU requires minimum balance (prevents instant suspension)
- **Debt tolerance**: Balance can go negative, service continues, time to add funds
- **Resource caps**: Global vCPU/instance limits prevent overprovisioning
- **Per-product visibility**: Identify and adjust high-cost products
- **Account suspension**: When balance < 0, withdrawals blocked but services continue

#### 4. Efficient Rate Updates

**Period-boundary rate changes** ensure accurate billing and rate stability within periods.

When admin updates a SKU rate:
```solidity
_updateSKU(skuId, name, runningRate, stoppedRate, vcpus, minimumDeposit)
```

- Stores new rate in `pendingSKUs[skuId]` with `effectivePeriod = nextPeriod`
- Gas cost: O(1) - no customer iteration
- Applied lazily when each customer's SKU is accessed
- Guarantees rates are constant within any given period

**Benefits**:
- Admin pays O(1) gas regardless of customer count
- Customers never see retroactive rate changes mid-period
- Accurate period charge calculations
- Cost amortized across customer operations

**New SKUs** apply immediately via:
```solidity
_addSKU(skuId, name, runningRate, stoppedRate, vcpus, minimumDeposit)
```

### Two Billing Models

#### ComputeBilling (Time-Based)

**For**: Continuous services (VMs, containers, persistent apps)

**How it works**:
1. **Accrual**: Charges accumulate continuously based on time × rate
   - Running apps: charged at `runningRate` tokens/second
   - Stopped apps: charged at `stoppedRate` tokens/second (lower)
2. **Settlement**: At period end, settler calls `settlePeriod(account, period)`
3. **Charge**: Deducts accumulated charges from BillingCore balance

**Features**:
- Running/stopped rate differential
- Resource tracking (vCPU and VM instance caps)
- Minimum deposit requirements per SKU
- State transitions update rates automatically
- Batch settlement for efficiency

**Functions**:
- `accrueCharges(account)` - Update accrued charges to current time
- `settlePeriod(account, period)` - Settle a specific period
- `settlePeriodBatch(accounts[], period)` - Batch settle multiple accounts
- `getOutstandingCharges(account)` - View accrued but unsettled charges
- `estimateCurrentPeriodCharges(account)` - Estimate current period

#### UsageBilling (Metered)

**For**: Event-driven services (API calls, data transfer, function executions)

**How it works**:
1. **Recording**: Admin calls `recordUsage(account, amount)` as events occur
   - Accumulates in `periodUsage[account][period]`
2. **Settlement**: At period end, admin calls `settlePeriod(account, period)`
3. **Charge**: Deducts recorded usage from BillingCore balance

**Features**:
- Record usage for current or specific periods
- Cannot modify usage after period is settled
- Batch recording and settlement
- Track unsettled usage across periods

**Functions**:
- `recordUsage(account, amount)` - Record usage in current period
- `recordUsageForPeriod(account, amount, period)` - Record for specific period
- `recordUsageBatch(accounts[], amounts[], period)` - Batch record
- `settlePeriod(account, period)` - Settle a specific period
- `settlePeriodBatch(accounts[], period)` - Batch settle
- `getUsageForPeriod(account, period)` - View usage details
- `getUnsettledUsage(account)` - Total unsettled across all periods

**Key difference**: Both use period settlement, but ComputeBilling accrues continuously based on time, while UsageBilling records discrete events.

### Period Management

**Periods** are calendar months since genesis:
- Period 0 = month of contract deployment
- Period 1 = following month
- etc.

**Period functions**:
```solidity
getCurrentPeriod() → uint40              // Current period number
getPeriodForTimestamp(timestamp) → uint40  // Period for specific time
getPeriodBounds(period) → (start, end)     // Timestamps for period
```

**Settlement timing**:
- Periods can only be settled AFTER they end (period < currentPeriod)
- Typically settled at the beginning of the next period
- Batch operations enable efficient end-of-month settlement

### Account Suspension

**Automatic suspension:**
- Triggered when balance < 0
- Blocks withdrawals only
- Services continue charging (debt increases)
- Auto-resumes when balance >= 0 (via deposit)

**No service interruption**: Even with negative balance, apps continue running and accruing charges. This provides customers time to add funds without service disruption.

## Key Functions

### For Customers

**Check balance:**
```solidity
getBalance(account) → int96               // Current balance (can be negative)
isAccountSuspended(account) → bool        // True if balance < 0
getAccount(account) → (balance, totalSpent, suspended)
```

**View charges:**
```solidity
// BillingCore - queries all products
getChargesForPeriod(account, period) → ProductCharges[]

// ComputeBilling
getOutstandingCharges(account) → uint96
getChargesForPeriod(account, period) → (amount, settled)
estimateCurrentPeriodCharges(account) → uint96

// UsageBilling
getUsageForPeriod(account, period) → (amount, settled, settledAt)
getUnsettledUsage(account) → uint96
```

**Manage balance:**
```solidity
deposit(amount)                   // Deposit to your account
depositFor(account, amount)       // Deposit to another account
withdraw(amount)                  // Withdraw (blocked if suspended)
```

### For Product Admins

**Manage SKUs (ComputeBilling):**
```solidity
// Add new SKU (applies immediately)
_addSKU(skuId, name, runningRate, stoppedRate, vcpus, minimumDeposit)

// Update existing SKU (applies next period)
_updateSKU(skuId, name, runningRate, stoppedRate, vcpus, minimumDeposit)

// Set resource limits
_setResourceCap(vcpuCap, vmInstanceCap)
```

**Record usage (UsageBilling):**
```solidity
recordUsage(account, amount)                           // Current period
recordUsageForPeriod(account, amount, period)          // Specific period
recordUsageBatch(accounts[], amounts[], period)        // Batch
```

**Settle periods:**
```solidity
// ComputeBilling
settlePeriod(account, period)
settlePeriodBatch(accounts[], period)

// UsageBilling
settlePeriod(account, period)
settlePeriodBatch(accounts[], period)
```

### For Platform Admins

**Manage products:**
```solidity
registerProduct(name, module) → productId    // Register new product module
setRevenueRecipient(productId, recipient)    // Set where revenue goes
withdrawRevenue(productId)                   // Claim revenue
deactivateProduct(productId)                 // Deactivate product
```

**Authorize settlers/reporters:**
```solidity
_setSettler(address, authorized)         // ComputeBilling
_setUsageReporter(address, authorized)   // UsageBilling
```

## Production Considerations

### Settlement Automation

Period settlement should be automated with a keeper/cron job:

```solidity
// End-of-month settlement
contract BillingSettler {
    function settleLastPeriod() external {
        uint40 lastPeriod = billingCore.getCurrentPeriod() - 1;

        // Get active accounts (from indexer/database)
        address[] memory computeAccounts = getActiveComputeAccounts();
        address[] memory usageAccounts = getActiveUsageAccounts();

        // Batch settle
        computeBilling.settlePeriodBatch(computeAccounts, lastPeriod);
        usageBilling.settlePeriodBatch(usageAccounts, lastPeriod);
    }
}
```

### Event-Based Analytics

All charges emit `PeriodCharge` events with full attribution:
```solidity
event PeriodCharge(
    address indexed account,
    uint8 indexed productId,
    uint40 indexed period,
    uint96 amount,
    int96 newBalance,
    bool suspended
);
```

These events enable rich analytics via subgraph indexing:
- Historical spending by period
- Per-product breakdown
- Account balance over time
- Suspension events

### Gas Optimization

**Batch operations**: Settle multiple accounts in single transaction
**Lazy evaluation**: SKU changes applied only when accessed
**Event-based reporting**: No complex on-chain storage for period tracking
**Simple charge logic**: Minimal storage operations per settlement
