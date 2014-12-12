using System;
using System.Runtime.InteropServices;
using System.Diagnostics;

namespace getForegroundWindowName
{
	class MainClass
	{
		[DllImport ("user32.dll")]
		public static extern IntPtr GetWindowThreadProcessId (IntPtr hWnd, out uint ProcessId);

		[DllImport ("user32.dll")]
		private static extern IntPtr GetForegroundWindow ();

		public static void Main (string[] args)
		{
			try {
				string name = getForegroundWindowName.MainClass.GetActiveProcessFileName ();
				System.Console.WriteLine (name);
			} catch {
			}
		}

		private static string GetActiveProcessFileName ()
		{
			IntPtr hwnd = GetForegroundWindow ();
			uint pid;
			GetWindowThreadProcessId (hwnd, out pid);
			Process p = Process.GetProcessById ((int)pid);
			return p.MainModule.FileName;
		}
	}
}