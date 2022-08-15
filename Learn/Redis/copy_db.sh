# set connection data accordingly
# 复制数据
source_host=
source_port=
source_db=
source_password=

target_host=
target_port=
target_db=
target_password=

# copy all keys without preserving ttl!
redis-cli -h $source_host -p $source_port -a $source_password -n $source_db keys \* | while read key; do
  echo "Copying $key"
  redis-cli --raw -h $source_host -p $source_port -a $source_password -n $source_db DUMP "$key" |
    head -c -1 |
    redis-cli -x -h $target_host -p $target_port -a $target_password -n $target_db RESTORE "$key" 0
done
