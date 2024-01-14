package model

type Strategy func(knapsack *Knapsack, items []Item) error
