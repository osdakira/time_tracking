#import "MBAAppDelegate.h"

@implementation MBAAppDelegate {
    NSTask *task;
}

- (void)awakeFromNib {
  _statusItem = [[NSStatusBar systemStatusBar] statusItemWithLength:NSVariableStatusItemLength];
  
  NSImage *menuIcon       = [NSImage imageNamed:@"Menu Icon"];
  NSImage *highlightIcon  = [NSImage imageNamed:@"Menu Icon"]; // Yes, we're using the exact same image asset.
  [highlightIcon setTemplate:YES]; // Allows the correct highlighting of the icon when the menu is clicked.

  [[self statusItem] setImage:menuIcon];
  [[self statusItem] setAlternateImage:highlightIcon];
  [[self statusItem] setMenu:[self menu]];
  [[self statusItem] setHighlightMode:YES];

  [self startLog];
}

- (void)applicationWillTerminate:(NSNotification *)application
{
    NSLog(@"Application will terminate.");
    [task terminate];
}

- (IBAction)sendReport:(id)sender {
    NSString *directory = [self getCurrentDirectoryPath];
    NSString *reportFile = [directory stringByAppendingString:@"/tracking.csv"];
}

- (IBAction)openDirectory:(id)sender {
    NSString *directory = [self getCurrentDirectoryPath];
    [[NSWorkspace sharedWorkspace]openFile:directory withApplication:@"Finder"];
}

- (NSString *) getCurrentDirectoryPath {
    NSString * current = [[NSFileManager defaultManager] currentDirectoryPath];
    return current;
}

- (NSString *) getTargetDirectoryPath {
    NSString *timeTrackingDirPath = @"/Users/osada/projects/time_tracking/mac";
    return timeTrackingDirPath;
}

- (void) startLog {
    NSString *directory = [self getTargetDirectoryPath];
    NSString *path = [directory stringByAppendingString:@"/time_tracking"];
    NSArray *args = [NSArray new];
    task = [NSTask launchedTaskWithLaunchPath:path arguments:args];
}
@end