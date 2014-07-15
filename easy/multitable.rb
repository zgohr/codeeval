count = 0
(1..12).each do |x|
    count += 1
    (1..12).each { |n| print (n*count).to_s.rjust(4) }
    puts ''
end

