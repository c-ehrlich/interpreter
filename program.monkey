let fizz = fn(x){
    if (x % 15 == 0) { return "fizzbuzz"; }
    if (x % 3 == 0) { return "fizz"; }
    if (x % 5 == 0) { return "buzz"; }
    return x;
};

let i = 1;

while (i <= 100) {
    puts(fizz(i));
    ++i;
}