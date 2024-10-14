lw 0 1 neg1
lw 0 2 ten
lw 0 3 one
noop
loop add 2 1 2
noop
noop
beq 2 0 done
noop
noop
noop
beq 0 0 loop
noop
noop
noop
done halt
neg1 .fill -1
ten .fill 10
one .fill 1
