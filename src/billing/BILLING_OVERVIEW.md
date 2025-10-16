# Billing System Overview

## What Is It?

A prepaid billing system for on-chain services. Accounts deposit tokens upfront, and services charge against those balances in real-time. Think of it like a prepaid phone plan, but for compute resources and API usage.

## Key Concepts

### Accounts

An **account** is any blockchain address that uses your services. This could be:
- A user's wallet
- A smart contract
- A multisig wallet
- Another protocol

Accounts deposit tokens to prepay for services. As they use services, their balance decreases. They can top up anytime.

### Products

A **product** is any service that charges for usage. Examples:
- Virtual machines
- API endpoints
- Data storage
- Network bandwidth

Each product is registered in the billing system and can charge accounts independently.

### Billing Periods

The system uses **calendar months** as billing periods, starting from the genesis timestamp when the contract is deployed.

- **Period 0**: The month when the contract is deployed
- **Period 1**: The next calendar month
- **Period 2**: The month after that, and so on

This natural calendar alignment makes accounting and reconciliation straightforward. All charges are accurately attributed to the period when the usage occurred, even if settlement happens later.

### Balance vs Debt

Every account has a **balance** (prepaid tokens) and potentially **debt** (owed tokens).

**When balance is sufficient:**
- Charges deduct from balance immediately
- Account continues operating normally

**When balance runs out:**
- Service doesn't stop immediately
- Charges create debt instead
- Account becomes suspended (can't withdraw remaining tokens)
- Services can continue charging the suspended account

**Paying off debt:**
- Anyone can pay an account's debt (not just the account owner)
- When debt reaches zero, the account automatically unsuspends
- Account can withdraw tokens again

This prevents service interruption while ensuring payment eventually happens.

## Two Billing Models

### 1. Compute Billing (Time-Based)

**Best for:** Continuous services that run over time
- Virtual machines
- Containers
- Persistent applications
- Long-running processes

**How it works:**
- Each service instance (app) is assigned a pricing tier (SKU)
- Apps can be in two states: **running** or **stopped**
- Each state has a different price per second
- The longer something runs, the more it costs
- Charges are attributed to the correct calendar month, even when settled later

**SKU System:**
SKUs (Stock Keeping Units) define pricing tiers with:
- Running rate (tokens per second when active)
- Stopped rate (tokens per second when paused)
- Resource allocation (vCPUs)
- Description

**Example SKUs:**
- Small VM: 0.01 tokens/hour running, 0.001 tokens/hour stopped, 2 vCPUs
- Medium VM: 0.05 tokens/hour running, 0.005 tokens/hour stopped, 4 vCPUs
- Large VM: 0.20 tokens/hour running, 0.020 tokens/hour stopped, 8 vCPUs

**State changes:**
- Start an app → Begins charging at running rate
- Stop an app → Switches to stopped rate (not zero!)
- Restart an app → Switches back to running rate
- Terminate an app → Stops all charges, final bill settled

**Why stopped isn't free:**
Even stopped services consume resources (storage, reserved capacity). The stopped rate reflects this.

**Resource Management:**
- System tracks global vCPU capacity
- Prevents overprovisioning
- Apps can change SKUs (upgrade/downgrade) with resource checks
- Capacity limits are configurable by administrators

**Account management:**
- One account can pay for multiple apps
- Apps can be reassigned to different accounts
- All apps for an account are billed together
- Efficient storage using packed structs (one slot per account)

**Settlement Optimization:**
The system intelligently handles settlements:
- Same period: Direct calculation
- One month boundary: Single boundary calculation
- Multiple months: Efficient period-by-period attribution

### 2. Usage Billing (Metered)

**Best for:** Event-driven services with variable usage
- API calls
- Data transfer
- Function executions
- Transaction processing

**How it works:**
- Service tracks usage amounts for each account per period
- Usage can be recorded as it happens (real-time) or in batches
- Charges accumulate until settlement
- Settlement happens after the period ends
- Each period's usage is tracked and settled independently

**Example pricing:**
- API endpoint: 0.001 tokens per call
- Data transfer: 0.10 tokens per GB
- AI inference: 0.05 tokens per request

**Recording Usage:**
Two approaches supported:
1. **Incremental**: Record usage as it happens
2. **Batch**: Report total usage for multiple accounts at once

**Settlement Process:**
- Can only settle completed periods (not the current one)
- Settlement charges the account and marks the period as finalized
- Supports batch settlement for efficiency
- Once settled, a period cannot be modified

**Period Attribution:**
All usage is attributed to the specific calendar month when it occurred:
- January usage is billed to Period N
- February usage is billed to Period N+1
- Accurate even if settlement happens months later

## Billing Lifecycle

### For Compute Services

**1. Account Setup**
- Account deposits tokens
- Billing data initialized on first app

**2. Service Launch**
- User requests a VM/app/container
- System assigns a pricing tier (SKU)
- Checks resource availability
- Billing starts immediately at the running rate

**3. During Service Life**
- Charges accumulate automatically over time
- Accurate period attribution as months change
- Account balance decreases continuously

**4. State Changes**
- User stops the service → Rate switches to stopped rate
- User restarts → Rate switches back to running rate
- User changes SKU → New rate applies after settlement

**5. Service End**
- User terminates the service
- Final charges calculated and settled
- Resources freed, billing stops

### For Usage Services

**1. Account Setup**
- Account deposits tokens
- Service ready to track usage

**2. Usage Recording**
- Service records usage to specific periods
- Each usage event updates the period total
- No immediate charges

**3. Settlement**
- After period ends (next month starts)
- All accumulated usage charged at once
- Period marked as settled
- Cannot modify settled periods

**4. Ongoing**
- New period starts automatically
- Usage continues in new period
- Previous periods can be settled anytime

## What Happens When Balance Runs Out?

**Immediate effects:**
- New charges create debt instead of deducting from balance
- Account marked as suspended
- Account cannot withdraw any remaining tokens

**Services continue:**
- For compute: Apps keep running and charging
- For usage: Usage keeps being recorded
- This prevents service disruption

**Recovery:**
- Anyone deposits tokens to pay the debt
- When debt reaches zero, account unsuspends automatically
- Account can withdraw again

**Why this design?**
- Prevents sudden service outages
- Gives grace period for payment
- Protects service providers (charges still recorded)
- Flexible payment (anyone can pay)

## Spending Visibility

The system tracks spending by calendar month:
- See total spending across all services for current month
- Break down spending by product
- View historical spending for any past month
- Useful for budgeting and cost management

Query capabilities:
- Current period spending for an account
- Detailed breakdown by product for any period
- Historical data preserved indefinitely

## Revenue for Service Providers

Each product (service) tracks its total revenue from all accounts. Service providers can:
- View accumulated revenue
- Withdraw revenue to their wallet
- Revenue includes partial payments from suspended accounts

## Technical Architecture

### Core Components

**BillingCore**: Central billing hub
- Manages accounts and balances
- Handles deposits and withdrawals
- Tracks products and spending
- Provides period calculation for entire system

**ComputeBilling**: Time-based billing module
- Manages SKUs and resource allocation
- Tracks running/stopped states
- Handles app lifecycle
- Optimized settlement across period boundaries

**UsageBilling**: Consumption-based billing module
- Records usage per period
- Manages batch operations
- Ensures period finality
- Direct account-based tracking

### Storage Optimization

The system uses efficient storage patterns:
- Packed structs to minimize storage slots
- Single slot per account for billing data
- Custom errors instead of strings
- Optimized period calculations

### Gas Efficiency

Several optimizations reduce transaction costs:
- Batch operations for multiple accounts
- Efficient period boundary calculations
- Minimal storage operations
- Reusable period calculation logic

## Common Scenarios

### Scenario 1: Running a VM

1. Alice deposits 1000 tokens
2. Alice launches a small VM (0.01 tokens/hour)
3. VM runs from January 15-25 (240 hours) → Costs 2.4 tokens
4. Alice stops the VM
5. Stopped VM continues through February (720 hours) → Costs 0.72 tokens
6. Alice terminates the VM in early March
7. January charges: 2.4 tokens (attributed to Period 0)
8. February charges: 0.72 tokens (attributed to Period 1)
9. Final balance: 996.88 tokens

### Scenario 2: Using an API

1. Bob deposits 500 tokens
2. Bob's app makes API calls throughout January
3. Each call costs 0.001 tokens, tracked for Period 0
4. End of January: 1000 calls recorded
5. February 1st: January usage settled → 1 token charged
6. Bob's balance: 499 tokens
7. February usage tracked separately in Period 1

### Scenario 3: Running Out of Balance

1. Carol has 5 tokens
2. Carol's VM costs 1 token/hour
3. After 5 hours, balance is 0
4. VM keeps running, creates 1 token of debt per hour
5. After 3 more hours, Carol has 3 tokens of debt
6. Carol's account is suspended (but VM keeps running)
7. Carol deposits 10 tokens
8. 3 tokens pay off debt, account unsuspends
9. New balance: 7 tokens

### Scenario 4: Company Account

1. Startup runs 10 apps across 3 accounts
2. Each account pays for its own apps
3. Apps can be reassigned between accounts (both settle first)
4. CFO views spending breakdown per account per month
5. Each account funded by different departments
6. Accurate monthly reports align with calendar months

### Scenario 5: Cross-Period Settlement

1. Dave's service runs continuously from January 20 to February 10
2. Settlement happens on February 15
3. System automatically:
   - Calculates January 20-31 charges → Attributes to January period
   - Calculates February 1-10 charges → Attributes to February period
   - Each period gets accurate allocation
4. Historical records show correct monthly spending

## Benefits

**For Service Users:**
- Predictable prepaid billing
- Real-time balance tracking
- Natural calendar month alignment
- No surprise bills
- Grace period if balance runs out
- Accurate historical records

**For Service Providers:**
- Guaranteed payment upfront
- Automatic charge collection
- Revenue tracking per product
- Protection against non-payment
- Flexible pricing models (time or usage)
- Efficient gas costs

**For Both:**
- Transparent on-chain pricing
- Full auditability
- Programmable (smart contracts can be accounts)
- Composable with other DeFi systems
- Calendar-aligned accounting
- Optimized for gas efficiency