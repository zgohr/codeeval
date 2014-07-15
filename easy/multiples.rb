File.readlines(ARGV[0]).each do |l|
    l.strip!
    x, n = l.split(',')
    i = count = 0
    while i < x.to_i
        count += 1
        i = n.to_i * count
    end
    puts i
end
