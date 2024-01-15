package model

import "github.com/shopspring/decimal"

// ValueType the type used for representing item values and costs, and also the knapsack capacity. We are using
// Decimals here to avoid issues with floating-point arithmetic. However, it may make sense to use float64 for
// better performance, at the cost of having slightly less precision.
type ValueType = decimal.Decimal
