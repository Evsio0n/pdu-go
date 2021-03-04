# DIS PDU Library for Golang

## Current standard: IEEE 1278.1

Currently: Dirty and CANNOT USE.

Referencing: http://faculty.nps.edu/brutzman/vrtp/mil/navy/nps/disEnumerations/JdbeHtmlFiles/pdu/29.htm https://www.sisostds.org/DesktopModules/Bring2mind/DMX/API/Entries/Download?Command=Core_Download&EntryId=52359&PortalId=0&TabId=105


PDU Components:

Item Name|Bit Length|Opt|Opt Ctl|Rpt|Rpt Ctl
-|-|-|-|-
PDU Header Record	|96
Entity ID Record	|48				
Force ID Field	|8
'#of Articulation Parameters (n) Field	'|8
Entity Type Record	|64				
Alternative Entity Type Record	|64				
Entity Linear Velocity Record	|96				
Entity Orientation Record	|96				
Entity Location Record	|192				
Entity Appearance Record	|32				
Dead Reckoning Parameters Record	|320				
Entity Marking record	|96				
Entity Capabilities Record	|32				
Articulation Parameter Record	|128	|Yes	|Number of Articulation Parameters	|Yes|	Number of Articulation Parameters