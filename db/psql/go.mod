module github.com/squishedfox/fictional-fiesta/db/psql

go 1.24.3

replace github.com/squishedfox/fictional-fiesta/db => ../

require github.com/squishedfox/fictional-fiesta/db v0.0.0-20250716190300-d695f25eb37e

require github.com/lib/pq v1.10.9 // indirect
