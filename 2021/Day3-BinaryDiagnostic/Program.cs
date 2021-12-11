using System;

namespace Day3_BinaryDiagnostic
{
    class Program
    {
        public string[] readFile(string path)
        {
            string[] lines = System.IO.File.ReadAllLines(path);
            return lines;
        }

        public int[] sumCols(string[] lines)
        {
            int[] sumDiag = new int[lines[0].Length];

            foreach(string line in lines)
            {
                int i = 0; 
                var lineChars = line.ToCharArray();
                
                foreach( char c in lineChars)
                {
                    var value =Int32.Parse(c.ToString());
                    sumDiag[i] = sumDiag[i] + value;
                    i++;
                }
            }
            return sumDiag;
        }

        public int gamaRate(int[] code)
        {
            string average = "";

            foreach(int number in code)
            {
                if( number > 500 )
                {
                    average = average + "1";
                }
                else
                {
                    average = average + "0";
                }               
            }
            Int64.Parse(average);
            var rate = Convert.ToInt32(average, 2);
            return rate;
        }

        public int epsilonRate(int[] code)
        {
            string average = "";

            foreach(int number in code)
            {
                if( number < 500 )
                {
                    average = average + "1";
                }
                else
                {
                    average = average + "0";
                }               
            }
            Int64.Parse(average);
            var rate = Convert.ToInt32(average, 2);
            return rate;
        }

        static void Main(string[] args)
        {
            Program p = new Program();

            var lines = p.readFile("Diagnostics.txt");

            var code = p.sumCols(lines);

            var gamma = p.gamaRate(code);

            var epsilon = p.epsilonRate(code);

            var consumption = gamma * epsilon;

            Console.WriteLine("Power Consumption: " + consumption);
        }
    }
}
