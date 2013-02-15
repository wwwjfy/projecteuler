local math = require 'math'

local primes = {2, 3}
local i = 5

while true do
    sqrt_i = math.sqrt(i)
    local flag = false
    for j, v in ipairs(primes) do
        if sqrt_i < v then
            flag = true
            table.insert(primes, i)
            break
        elseif i % v == 0 then
            flag = true
            break
        end
    end
    if not flag then
        table.insert(primes, i)
    end

    i = i + 2

    if i > 2000000 then
        break
    end
end

local final = 0
for j,v in ipairs(primes) do
    final = final + v
end

print(final)