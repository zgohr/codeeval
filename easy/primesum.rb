total = count = 0
i = 2
while count < 1000
    broken = false
    (2..i/2).each do |n|
        if i % n == 0
            broken = true
            break
        end
    end
    if !broken
        count += 1
        total += i
    end
    i += 1
end
puts total
