File.readlines(ARGV[0]).each do |l|
    a, b = l.split(',')
    b.strip!
    idx = a.index(b)
    if idx and idx + b.length == a.length
        puts '1'
    else
        puts '0'
    end
end
