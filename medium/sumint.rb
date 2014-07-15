File.readlines(ARGV[0]).each do |l|
    l.strip!
    arr = l.split(',')
    max_sum = 0
    positive_sum = 0
    arr.each do |v|
        positive_sum = [v.to_i + positive_sum, 0].max
        max_sum = [max_sum, positive_sum].max
    end
    puts max_sum
end
