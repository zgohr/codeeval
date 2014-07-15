File.readlines(ARGV[0]).each do |l|
    l.strip!
    a, b, n = l.split
    out = []
    (1..n.to_i).each do |x|
        if x % a.to_i == 0 and x % b.to_i == 0
            out.push('FB')
        elsif x % a.to_i == 0
            out.push('F')
        elsif x % b.to_i == 0
            out.push('B')
        else
            out.push(x.to_s)
        end
    end
    puts out.join(' ')
end
