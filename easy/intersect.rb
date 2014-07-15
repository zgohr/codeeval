File.readlines(ARGV[0]).each do |l|
    l.strip!
    a, b = l.split(';').collect{|x| x.split(',')}
    puts (a & b).join(',')
end
