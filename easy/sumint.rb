s = 0
File.readlines(ARGV[0]).each do |l|
    s += l.strip.to_i
end
puts s.to_s
