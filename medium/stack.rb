class Stack
  def initialize
    @e = []
  end

  def push(e)
    @e.push(e)
  end

  def pop
    @e.pop
  end
end

File.readlines(ARGV[0]).each do |l|
  the_stack = Stack.new
  l.split.each { |i| the_stack.push(i) }
  result = []

  loop do
    break if (e = the_stack.pop).nil?
    result.push(e)
    the_stack.pop
  end
  puts result.join(' ')
end
