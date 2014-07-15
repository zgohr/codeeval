File.readlines(ARGV[0]).each do |l|
    puts l.strip.split(',').uniq.join(',')
end
