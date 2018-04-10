# hammer
Iterative frontend to vegeta
```
./hammer -url http://4.3.2.1/img/redapple002.jpg -start 10000 -duration 5 -increment 2000 

# Trying rate 10000, iteration 2000
+ rate 10000 success 100.0% errs 0 99th percentile: 3.271375ms

# Trying rate 12000, iteration 2000
+ rate 12000 success 100.0% errs 0 99th percentile: 1.60514ms

# Trying rate 14000, iteration 2000
+ rate 14000 success 100.0% errs 0 99th percentile: 1.998847ms

# Trying rate 16000, iteration 2000
+ rate 16000 success 100.0% errs 0 99th percentile: 2.053279ms

# Trying rate 18000, iteration 2000
+ rate 18000 success 100.0% errs 0 99th percentile: 2.566002ms

# Trying rate 20000, iteration 2000
+ rate 20000 success 100.0% errs 0 99th percentile: 44.733598ms

# Trying rate 22000, iteration 2000
- Requests 110000, ok 56261  51.1%
  Dropping rate to 21000 and increment by 1000
  Pausing 3m0s

# Trying rate 21000, iteration 1000
- Requests 105000, ok 37633  35.8%
  Dropping rate to 20500 and increment by 500
  Pausing 4m0s

# Trying rate 20500, iteration 500
- Requests 102500, ok 37540  36.6%
  Dropping rate to 20250 and increment by 250
  Pausing 5m0s

# Trying rate 20250, iteration 250
- Requests 101250, ok 31411  31.0%
  Dropping rate to 20125 and increment by 125

* rate 20000 success 100.0% errs 0 99th percentile: 44.733598ms
best run 20000
```
