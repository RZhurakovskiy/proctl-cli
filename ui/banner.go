package ui

import "fmt"

func ShowBanner() {
	fmt.Println(`                            __  .__   
_____________  ____   _____/  |_|  |  
\____ \_  __ \/  _ \_/ ___\   __\  |  
|  |_> >  | \(  <_> )  \___|  | |  |__
|   __/|__|   \____/ \___  >__| |____/
|__|                     \/           `)
	fmt.Println(`========================================
          proctl v1.1
   Process Control Utility (Go)
   Автор: Roman Zhurakovskiy
========================================`)
}
