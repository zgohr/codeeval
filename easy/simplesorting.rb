File.readlines(ARGV[0]).each do |l|
    input = l.strip!.split.map { |i| i.to_f }
    puts input.sort.map{ |i| "%0.3f" % i }.join(' ')
end
