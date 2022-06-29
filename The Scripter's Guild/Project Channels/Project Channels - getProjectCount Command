{{ $infiniteProjectCount := toInt (dbGet 1 "infiniteProjectCount").Value }}
{{ $legacyHaloProjectCount := toInt (dbGet 1 "legacyHaloProjectCount").Value }}
{{ $unrealProjectCount := toInt (dbGet 1 "unrealProjectCount").Value }}
{{ $unityProjectCount := toInt (dbGet 1 "unityProjectCount").Value }}
{{ $coreProjectCount := toInt (dbGet 1 "coreProjectCount").Value }}
{{ $fortniteProjectCount := toInt (dbGet 1 "fortniteProjectCount").Value }}

{{ print 
"Infinite: " $infiniteProjectCount " | "
"Legacy Halo: " $legacyHaloProjectCount " | "
"Unreal: " $unrealProjectCount " | "
"Unity: " $unityProjectCount " | "
"Core: " $coreProjectCount " | "
"Fortnite: " $fortniteProjectCount 
}}
