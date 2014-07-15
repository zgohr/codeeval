File.readlines(ARGV[0]).each do |l|
  total = l.strip.to_i
  n = 0

  [5, 3, 1].each do |c| 
    div = total/c
    n = n + div
    total = total - (div * c)
  end

  puts n
end
