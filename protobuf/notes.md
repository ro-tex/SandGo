## Notes
 - no inheritance between message types
 - ids under 16 take one less byte to encode, so use those for fields you encounter often
 - repeated items encode their id with each item, so make sure they use an id under 16 as well
 
