# procgen

This is a sandbox repo for playing with procedural generation

## Attempt 1

The idea is to generate/emulate a medieval map (with castles, towns, market,
farms, etc.) using a cellular automata.

The entry point is a simple house, which triggers building of a farm nearby.
The farm might trigger building of a windmill. Several houses will require a
market and later can be converted into a castle.

Some TODO-list:
- consider threating some cells as a single cell: for example, a set of nearby
  houses might be considered as a village and should be handled as a single
  entity
- consider adding some cost models, i.e. base cell type selection not on random
  and surrounding cells, but on needs: a house needs food, so we increase need
  for farm. If there are several houses, more farms or even a windmill is
  generated. Due to some circumstances a house might be destroyed: due to lack
  of food or because of being overcrowded
- [VISUALS] add icons instead of colored rectangles
- [VISUALS] consider adding sub-types for cells to make the resulting picture
  more interesting
