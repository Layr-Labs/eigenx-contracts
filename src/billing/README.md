# Billing System Overview

## Customer Assumptions

Customers expect the ability to:
1. **Predict costs**: Know if their balance will cover their workload across all products
2. **See per-product breakdown**: Understand charges per product by time period
3. **Control spending**: Adjust balance/resources when usage changes
4. **Withdraw funds**: Get unused funds out of the system

## Core Product Decisions

### USDC Only
The system uses USDC as the single payment token. Multi-currency support is out of scope.

### Anyone Can Pay for Any App
Anyone can deposit funds to pay for someone else's app via `depositFor(account, amount)`. To prevent misuse, we use **segmented accounts** - dedicated accounts that can only pay for specific apps and cannot withdraw funds.

This segmentation happens at the account level (not per individual resource) because tracking individual resources in a shared billing contract is too gas-intensive. The infrastructure for this is built but not yet enabled.

### Single Account Per User
One billing account pays for all products a user consumes (Compute, AI, etc.). This simplifies the experience - one balance to manage, one place to see all spending, easy cross-product cost analysis.

## System Architecture

### Hub-and-Spoke Design

**BillingCore (Hub):**
- Central source of truth for all accounts and balances
- Balances can go negative (debt tolerance)
- Manages product registration and withdrawal requests
- Calculates billing periods (calendar months from genesis)

**Product Modules (Spokes):**
- Each product (Compute, AI, etc.) has its own billing module
- Modules track usage and call BillingCore to charge accounts
- All use the same pattern: accumulate charges → settle at period end → charge account

### Unified Settlement Pattern
Both billing types follow the same monthly settlement cycle:
1. **Accumulate** - Track charges during the period
2. **Settle** - Calculate final amount at period end
3. **Charge** - Deduct from account balance in BillingCore

This provides predictable monthly billing across all products.

## Two Billing Types

### ComputeBilling (Time-Based)
For continuous services like VMs and containers. Charges accumulate based on time × rate. Apps have two states:
- **Running**: Higher rate when VM is active
- **Stopped**: Lower rate when VM is paused but resources reserved

Settlement happens at month end, converting accumulated time-based charges to a period charge.

### UsageBilling (Event-Based)
For metered services like API calls or data transfer. Usage events are recorded as they occur and accumulate within the period. Settlement happens at month end, summing all usage for the period.

## Withdrawing Funds

The withdrawal process ensures no pending charges appear after you withdraw:

### Phase 1: Request Withdrawal
- Must have zero active resources across all products
- Marks current period as your "final billing period"
- Prevents creation of new resources

### Phase 2: Wait Period
- Wait until next period begins (if requested in January, withdraw in February)
- Ensures your final period is complete and can be fully settled
- All products have time to report final usage

### Phase 3: Complete Withdrawal
- Verify final period is settled
- Confirm no resources were created after request
- Withdraw available balance

You can cancel a withdrawal request at any time to resume normal operations.

## Safety Features

- **Settlement protection**: Periods can only be settled once, and only after they end
- **Balance enforcement**: Usage-based services (offchain) must check balances before allowing usage. Compute services (onchain) will be suspended if payment defaults
- **Resource limits**: Global vCPU and VM instance caps prevent overprovisioning
- **Activity tracking**: Timestamps verify no unauthorized activity after withdrawal
- **Product deactivation**: Products can be sunset, preventing charges after deactivation

## Billing Periods

Periods are calendar months counted from system genesis:
- Period 0 = First month of deployment
- Period 1 = Second month
- Settlement occurs after each period ends

## For Developers

This document covers the product and architecture essentials. For technical details, review the proof-of-concept implementation in the source code. Note that this POC:
- Has not been tested and likely contains bugs
- Is intended to illustrate the solution architecture for discussion
- Should be treated as a directional reference, not production-ready code
