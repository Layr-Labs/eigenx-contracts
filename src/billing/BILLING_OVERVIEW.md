# Billing System Overview

## Customer Assumptions

Customers want to:

1. **Predict costs**: Know if their balance will cover their workload across all products
2. **See per-product breakdown**: Understand charges per product by time period
3. **Control spending**: Adjust balance/resources when usage changes

## Product Decisions

### 1. Payments in USDC Only
The system uses a single payment token (USDC) hardcoded at deployment. Adding multiple currencies is complex and out of scope.

### 2. Graceful Migration from Free Tier
Pre-billing apps can be migrated via `enableAppBilling()`:
- Admin enables billing per-app basis
- High-value customers: longer grace period or permanent exemption (set `skuID = 0`)
- Low-value customers: email notice, enable billing after 2-3 months

### 3. Anyone Can Pay for Any App
Anyone can deposit funds to pay for someone else's app via `depositFor(account, amount)`.

**Why segmented accounts matter**: Without restrictions, depositing to an account means the owner could:
- Withdraw your funds
- Use them to pay for other apps they own

**Solution - Segmented, restricted accounts**: Move apps to dedicated accounts that:
- Only pay for one specific app
- Cannot withdraw funds
- Cannot own other apps

This segmentation happens at the **account level** (not resource level) because tracking individual resources (apps, API keys, etc.) in the billing contract is too gas-intensive/complex.

**Implementation options**:
1. **Smart contract account** - lacks withdrawal and multi-app ownership functions
2. **Protocol-level restrictions** - built-in controls on certain account types

**Example flow**:
1. Alice deploys app, bills to herself: `createApp(skuID, aliceAddress)`
2. App uses Compute + AI → both charge Alice's account
3. Admin transfers app to restricted account: `enableAppBilling(app, skuID, restrictedAccount)`
4. All products now charge the restricted account
5. Community calls `depositFor(restrictedAccount, amount)` - funds locked to that app only

### 4. Single Account Model
One billing account pays for all products (Compute, AI, etc.). This is a product decision contingent on technical feasibility.

**Rationale**:
- Simplifies customer experience: one balance, one view of spending
- Enables cross-product cost analysis (see spending distribution)
- Single `getEffectiveBalance()` call shows position across all services
- Reduces cognitive load vs managing multiple accounts per product

**Technical enabler**: Hub-and-spoke architecture allows multiple product modules to charge the same account via `BillingCore.charge()`.

## System Architecture

### Hub-and-Spoke Design

**BillingCore (Hub):**
- Single source of truth for accounts and balances
- Each account has one `int96 balance` (positive = credit, negative = debt)
- Tracks all registered products (Compute = Product 1, AI = Product 2, etc.)
- Stores settled charges per account, per product, per period
- Provides period calculation (calendar months from genesis)

**Product Modules (Spokes):**
- `ComputeBilling`: Time-based - continuous accrual, auto-settlement
- `UsageBilling`: Metered - discrete events, admin-triggered settlement
- Each module calls `BillingCore.charge()` to deduct from account balance
- Multiple products can charge the same account

### Key Data Structures

**Account:**
```solidity
struct Account {
    int96 balance;        // Can go negative (debt)
    uint96 totalSpent;
    uint40 lastActivity;
    bool suspended;       // True when balance < 0
}
```

**Product:**
```solidity
struct Product {
    string name;
    address billingModule;      // ComputeBilling, UsageBilling, etc.
    address revenueRecipient;
    uint96 unclaimedRevenue;
    bool isActive;
}
```

**SKU (for Compute products):**
```solidity
struct SKU {
    uint96 runningRate;      // tokens/second
    uint96 stoppedRate;      // tokens/second
    uint16 vcpus;
    uint96 minimumDeposit;   // prevents instant suspension
    string description;
}
```

### How It Delivers on Customer Needs

#### 1. Predict Costs

**Effective Balance** = balance - outstanding charges across ALL products

```solidity
getEffectiveBalance(account) → {
    balance,
    outstandingCharges,    // sum from all products
    effectiveBalance,      // balance - outstanding
    willCoverCharges,
    breakdown[]            // per-product outstanding
}
```

BillingCore queries each product module for `getOutstandingCharges(account)`, sums them, and returns breakdown. Customer sees if effective balance covers pending charges before they settle.

#### 2. Per-Product Breakdown

**Settled charges** (historical, stored in BillingCore):
```solidity
getSettledChargesForPeriod(account, period) → (productIds[], names[], amounts[])
```

**Accrued charges** (includes unsettled, queries modules):
```solidity
getAccruedChargesForPeriod(account, period) → ProductCharges[] {
    productId, productName, amount
}
```

Both return arrays indexed by product. Customer can see: "I spent 80% on Compute (Product 1), 15% on AI (Product 2), etc last month."

**Calendar month periods**: Period 0 = deployment month, Period 1 = next month, etc. Natural accounting alignment.

#### 3. Control Spending

- **Prepaid model**: Deposit upfront, charges deduct automatically
- **Minimum deposits**: Each SKU requires minimum effective balance (prevents instant suspension)
- **Debt grace period**: Balance can go negative, service continues, time to add funds
- **Resource caps**: Global vCPU/instance limits prevent overprovisioning
- **Per-product visibility**: Identify and adjust high-cost products

#### 4. Efficient Rate Updates (Constant Gas)

**Problem**: Updating rates for all customers = O(n) gas where n = customer count

**Solution**: Pending rates + lazy evaluation

```solidity
setSKURate(skuID, ..., runningRate, stoppedRate, vcpus, minimumDeposit)
```

- Stores new rate in `pendingSKUs[skuID]` with `effectivePeriod = nextPeriod`
- Gas: O(1) regardless of customer count (no iteration)
- When customer's next settlement occurs, system checks if pending rate should apply
- Applied per-customer lazily, not upfront

Example:
- Admin updates SKU 1 rate at t=0 → 50k gas (constant)
- 1000 customers using SKU 1
- Each customer's next settlement applies new rate automatically
- Total cost amortized across customer operations, not paid upfront by admin

### Two Billing Models

#### ComputeBilling (Time-Based)

- **For**: Continuous services (VMs, containers, persistent apps)
- **Accrual**: Charges = rate × elapsed time since last settlement
- **Settlement**: Auto-triggered on app state changes (start, stop, terminate)
- **Cross-period**: Efficiently attributes charges to correct month even when settling across boundaries
- **User-triggered**: Settlement happens during app lifecycle operations

#### UsageBilling (Metered)

- **For**: Event-driven services (API calls, data transfer, function executions)
- **Recording**: Admin calls `recordUsage(account, amount, period)` as events occur
- **Settlement**: Admin calls `settlePeriod(account, period)` after period ends
- **Batch operations**: Can settle multiple accounts at once
- **Admin-triggered**: Explicit settlement by service provider

**Key difference**: ComputeBilling = continuous + user-triggered, UsageBilling = discrete + admin-triggered

### Account vs App Suspension

**Account suspension (automatic):**
- Triggered when balance < 0
- Blocks withdrawals only
- Services continue charging (debt increases)
- Auto-unsuspends when balance >= 0

**App SUSPENDED state (admin action):**
- Admin calls `suspendApp(app)` for prolonged non-payment
- Removes billing, frees resources
- App owner must call `unsuspendApp(app)` after paying debt
- Checks minimum deposit requirement before restarting

### Withdrawal Policy (TODO)

**Current issue**: Withdrawals are currently allowed while unsettled charges exist, particularly problematic for metered products (AI, usage-based services) where charges accrue but haven't been settled yet.

**Options under consideration**:
1. **Lock withdrawals** - Require users to unsubscribe from all products AND wait for final billing period settlement before allowing withdrawals
2. **Remove withdrawals completely** - Simpler implementation but carries UX/legal risks that may not be acceptable to customers who expect access to prepaid funds

## Key Functions

### For Customers

**Check financial position:**
```solidity
getEffectiveBalance(account) → EffectiveBalance
getBalance(account) → int96
isAccountSuspended(account) → bool
```

**View charges:**
```solidity
getAccruedChargesForPeriod(account, period) → ProductCharges[]  // includes unsettled
getSettledChargesForPeriod(account, period) → (ids[], names[], amounts[])  // historical
```

**Manage balance:**
```solidity
deposit(amount)
depositFor(account, amount)
withdraw(amount)
```

### For Admins

**Update rates (O(1) gas):**
```solidity
setSKURate(skuID, name, runningRate, stoppedRate, vcpus, minimumDeposit)
// Takes effect next period, applied lazily per customer
```

**Manage products:**
```solidity
registerProduct(name, module) → productId
setRevenueRecipient(productId, recipient)
withdrawRevenue(productId)
```

**Manage apps:**
```solidity
suspendApp(app)      // Stop billing, free resources
setResourceCap(vcpuCap, vmInstanceCap)
```
