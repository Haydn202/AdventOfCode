using System;

namespace CSharp
{
    class Program
    {
        public string[] readFile(string path)
        {
            string[] lines = System.IO.File.ReadAllLines(path);
            return lines;
        }

        public int[] convertToInt(string[] lines)
        {
            int[] readings = new int[lines.Length];
            int count = 0;

            foreach (string line in lines)
            {
                readings[count] = Int32.Parse(line);
                count++;
            }
            return readings;
        }

        public int[] getAvarages(int[] readings)
        {
            int[] avarages = new int[readings.Length - 2];
            for (int i = 0; i < readings.Length - 2; i++)
            {
                int aver = (readings[i] + readings[i + 1] + readings[i + 2]);
                avarages[i] = aver;
            }
            return avarages;
        }

        public int countDepths(int[] avarages)
        {
            int previousReading = avarages[0];
            int increasesInDepth = 0;
            
            foreach (int avarage in avarages)
            {
                int currentReading = avarage;

                if (currentReading > previousReading)
                {
                    increasesInDepth++;
                }
                previousReading = currentReading;
            }
            return increasesInDepth;
        }

        static void Main(string[] args)
        {
            Program p = new Program();

            var lines = p.readFile("../Day1-SpnarSweep/readings.txt");

            var readings = p.convertToInt(lines);

            var avarages = p.getAvarages(readings);

            Console.WriteLine(p.countDepths(avarages));
        }
    }
}
