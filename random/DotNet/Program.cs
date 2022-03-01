int n = 20;
int a = 0;
while (a < n)
{
    int b = n - a;
    while (b > 0)
    {
        Console.Write("*");
        b--;
    }
    Console.WriteLine();
    a++;
}
Console.WriteLine();
a = 0;
while(a < n) {
    int b = 0;
    while (b < a + 1) {
        Console.Write("*");
        b++;
    }
    Console.WriteLine();
    a++;
}
Console.WriteLine();
a = 0;
while (a < n)
{
    int b = 0;
    while (b < a)
    {
        Console.Write(" ");
        b++;
    }
    b = n - a;
    while (b > 0)
    {
        Console.Write("*");
        b--;
    }
    Console.WriteLine();
    a++;
}
Console.WriteLine();
a = 0;
while (a < n) {
    int b = a + 1;
    while (b < n) {
        Console.Write(" ");
        b++;
    }
    b = 0;
    while (b < a + 1) {
        Console.Write("*");
        b++;
    }
    Console.WriteLine();
    a++;
}