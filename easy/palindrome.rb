1000.downto(1) do |x|
    broke = false
    (2..x/2+1).each do |n|
        if x % n == 0
            broke = true
            break
        end
    end
    if not broke and x.to_s.reverse == x.to_s
        puts x
        break
    end
end
