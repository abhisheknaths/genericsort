heap-compare-bench:
	go test -bench BenchmarkCompareHeapSort -v -cpu=1 -count=10 -run=^X

heap-compare-mem:
	go test -bench BenchmarkCompareHeapSort -v -cpu=1 -benchmem -run=^X

merge-compare-bench:
	go test -bench BenchmarkCompareMergeSort -v -cpu=1 -count=10 -run=^X

merge-compare-mem:
	go test -bench BenchmarkCompareMergeSort -v -cpu=1 -benchmem -run=^X

build-bench-image:
	docker build -t benchmarksort:1 -f bench/Dockerfile .

docker-bench-heap:
	docker build -t benchmarksort:1 -f bench/Dockerfile .; \
	docker run --cpus=1 --memory 500MB --env BENCH_COMMAND="go test -bench BenchmarkCompareHeapSort -v -cpu=1 -count=10 -run=^X" benchmarksort:1

docker-bench-heap-mem:
	docker build -t benchmarksort:1 -f bench/Dockerfile .; \
	docker run --cpus=1 --memory 500MB --env BENCH_COMMAND="go test -bench BenchmarkCompareHeapSort -v -cpu=1 -benchmem -run=^X" benchmarksort:1

docker-bench-merge:
	docker build -t benchmarksort:1 -f bench/Dockerfile .; \
	docker run --cpus=1 --memory 500MB --env BENCH_COMMAND="go test -bench BenchmarkCompareMergeSort -v -cpu=1 -count=10 -run=^X" benchmarksort:1

docker-bench-merge-mem:
	docker build -t benchmarksort:1 -f bench/Dockerfile .; \
	docker run --cpus=1 --memory 500MB --env BENCH_COMMAND="go test -bench BenchmarkCompareMergeSort -v -cpu=1 -benchmem -run=^X" benchmarksort:1
