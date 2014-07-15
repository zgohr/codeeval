File.readlines(ARGV[0]).each do |l|
    puts l.strip.split('').collect{|s| s.to_i}.inject(:+)
end
