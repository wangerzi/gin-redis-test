# Redis Test By gin
You can deploy a simple web api in 1min to test redis connection and got redis information. 

Such as some docker environment that can't bin/sh directly (I mean AWS Fargate)

## Usage

Run server
```shell script
docker run -d -p 80:80 wj2015/gin-redis-test:v1.0.0
```

Access by web, url pattern: `/redis/:addr/:pwd/:db` to got all information.

```shell script
$ curl http://localhost/redis/127.0.0.1:6379/password/0
# if success:
{"data":{"info":{"Keys":[],"MemoryInfo":{"active_defrag_running":"0","allocator_active":"2342912","allocator_allocated":"1892600","allocator_frag_bytes":"450312","allocator_frag_ratio":"1.24","allocator_resident":"4800512","allocator_rss_bytes":"2457600","allocator_rss_ratio":"2.05","lazyfree_pending_objects":"0","maxmemory":"0","maxmemory_human":"0B","maxmemory_policy":"noeviction","mem_allocator":"jemalloc-5.1.0","mem_aof_buffer":"0","mem_clients_normal":"50958","mem_clients_slaves":"0","mem_fragmentation_bytes":"9285016","mem_fragmentation_ratio":"6.33","mem_not_counted_for_evict":"0","mem_replication_backlog":"0","number_of_cached_scripts":"1","rss_overhead_bytes":"6225920","rss_overhead_ratio":"2.30","total_system_memory":"1875775488","total_system_memory_human":"1.75G","used_memory":"1782432","used_memory_dataset":"916842","used_memory_dataset_perc":"93.59%","used_memory_human":"1.70M","used_memory_lua":"35840","used_memory_lua_human":"35.00K","used_memory_overhead":"865590","used_memory_peak":"2259504","used_memory_peak_human":"2.15M","used_memory_peak_perc":"78.89%","used_memory_rss":"11026432","used_memory_rss_human":"10.52M","used_memory_scripts":"216","used_memory_scripts_human":"216B","used_memory_startup":"802832"},"ServerInfo":{"arch_bits":"64","atomicvar_api":"atomic-builtin","config_file":"","configured_hz":"10","executable":"/data/redis-server","gcc_version":"8.3.0","hz":"10","lru_clock":"1894169","multiplexing_api":"epoll","os":"Linux 4.18.0-193.6.3.el8_2.x86_64 x86_64","process_id":"1","redis_build_id":"4935af324665042b","redis_git_dirty":"0","redis_git_sha1":"00000000","redis_mode":"standalone","redis_version":"6.0.1","run_id":"cccaa712cc3500efc7fefd910883057d66d21ff1","tcp_port":"6379","uptime_in_days":"30","uptime_in_seconds":"2670991"}}},"message":"success","status":true}  
# if failed:
{"data":{"keys":null},"message":"ERR AUTH \u003cpassword\u003e called without any password configured for the default user. Are you sure your configuration is correct?","status":false}
```

## Development
Environment: `GO111MODULE=on;GOPROXY=https://goproxy.cn,direct`

```shell script
$ go mod download
$ go build
$ ./gin-redis-test
```