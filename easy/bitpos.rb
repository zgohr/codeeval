File.readlines(ARGV[0]).each do |l|
    l.strip!
    n, p1, p2 = l.split(',')
    b = n.to_i.to_s(2)
    puts (b[-1 * p1.to_i] == b[-1 * p2.to_i]).to_s
end
