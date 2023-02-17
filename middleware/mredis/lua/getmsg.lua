local key = KEYS[1];
local list = redis.call('lrange',key,'0','-1');

if(#list > 0) then
    -- 清空list
    redis.call('ltrim',key,'1','0');
    return list;
end
return;
