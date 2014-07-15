File.readlines(ARGV[0]).each do |l|
    l.strip!
    nums, s = l.split(';')
    nums = nums.split(',')
    pairs = []
    nums.each_with_index do |n, i|
        (0..i-1).each do |i2|
            pairs.push("#{nums[i2]},#{n}") if n.to_i + nums[i2].to_i == s.to_i
        end
    end
    if pairs.length > 0
        puts pairs.reverse.join(';')
    else
        puts "NULL"
    end
end
