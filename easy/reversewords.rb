File.readlines(ARGV[0]).each do |l|
    l.strip!
    next if l == ''
    puts l.split.reverse.join(' ')
end
