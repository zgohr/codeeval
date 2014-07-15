class Integer
  def palindrome?
    return self.to_s == self.to_s.reverse
  end
end


File.readlines(ARGV[0]).each do |l|
    int = l.to_i
    iteration = 0
    until int.palindrome?
      int += int.to_s.reverse.to_i
      iteration += 1
    end
    puts "#{iteration} #{int}"
end
