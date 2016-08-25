1.Parallel Test
Run:
 	$ go test -bench=".*"
Result:
	PASS
	#Benchmark    function name	    CPU	   loop times	 each loop's time
	Benchmark_AmCertListServeParallel-4	     200000	         7011 ns/op
	Benchmark_AmCertListServe-4        	    1000000	         3906 ns/op
	#       run dir    total run time
	ok  	test		22.624s

2.Usual Test
Run:
	$ go test -v
