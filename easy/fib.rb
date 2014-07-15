
def fib(n)
    phi = ((1 + Math.sqrt(5)) / 2)
    return ((phi ** n - (-1 ** n / phi ** n)) / Math.sqrt(5)).to_i
end

File.readlines(ARGV[0]).each do |l|
    l.strip!
    puts fib(l.to_i)
end
