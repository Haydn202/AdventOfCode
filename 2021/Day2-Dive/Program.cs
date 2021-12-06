using System;
using System.Collections.Generic;

namespace Day1_Dive
{
    class Program
    {
        public string[] readFile(string path)
        {
            string[] lines = System.IO.File.ReadAllLines(path);
            return lines;
        }

        public List<KeyValuePair<string, int>> readDirections(string[] directions)
        {
            var headings = new List<KeyValuePair<string, int>>();

            foreach (string direction in directions)
            {
                string[] subs = direction.Split(' ');
                var heading = subs[0];
                var amplitude = Int32.Parse(subs[1]);   
                headings.Add(new KeyValuePair<string, int>(heading, amplitude));
            }
            return headings;
        }

        public string sumHeadings(List<KeyValuePair<string, int>> headings, int depth, int distance, int aim)
        {
            foreach (KeyValuePair<string, int> heading in headings)
            {
                switch (heading.Key)
                {
                    case "up":
                        aim = aim - heading.Value;
                        break;

                    case "down":
                        aim = aim + heading.Value;
                        break;

                    case "forward":
                        distance = distance + heading.Value;
                        depth = depth + heading.Value * aim;
                        break;
                }
            }
            return "Depth: " + depth + ", Distance: " + distance + ", Multiplied: " + depth * distance;
        }

        static void Main(string[] args)
        {
            Program p = new Program();

            var lines = p.readFile("../Day2-Dive/directions.txt");

            var directions = p.readDirections(lines);

            var heading = p.sumHeadings(directions, 0, 0, 0);

            Console.WriteLine(heading);
        }
    }
}
