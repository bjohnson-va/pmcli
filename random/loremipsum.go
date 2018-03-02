package random

var loremIpsumString = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut nec congue ante. Sed
at odio quis nisi ultricies aliquet vitae id metus. Donec augue nibh, mollis
eget suscipit ac, aliquam ut massa. Morbi congue tempus tempor. Cras sed lacinia
purus. Curabitur eu risus dolor. Donec dignissim risus orci, ut condimentum
libero elementum a. Donec eros turpis, porta a ex gravida, consequat tincidunt
neque. Mauris elit nibh, porta at eros imperdiet, malesuada ornare tellus. Nulla
facilisi. Suspendisse porta interdum nulla, eget feugiat orci cursus vitae. Duis
leo ligula, fringilla sed nibh ut, luctus ultrices dui. Pellentesque molestie
velit sem, sed ultrices mauris tincidunt nec.

Nullam suscipit dui vel placerat efficitur. Nulla facilisi. Ut lobortis lacinia
metus, a interdum risus. Quisque venenatis risus et placerat scelerisque. Morbi
turpis erat, varius placerat urna vitae, ullamcorper auctor sapien. Sed nibh
sapien, feugiat et vehicula sed, tincidunt in tortor. Morbi auctor euismod diam,
vel interdum metus tincidunt a.

Nam tempor, turpis et pellentesque suscipit, tortor justo vehicula arcu, ac
lacinia turpis ipsum sed nibh. Vivamus tortor odio, lacinia nec ultrices quis,
mattis sed lorem. Curabitur consequat mattis vulputate. Duis porttitor ex eget
erat ultricies, vel porta ante molestie. Cras dictum magna nulla, ac
sollicitudin quam fringilla id. Duis mollis molestie urna sit amet semper.
Aliquam nisi lectus, finibus id nunc vel, scelerisque aliquam mauris. Curabitur
non elit eget justo mollis volutpat suscipit at nunc. Ut velit massa, porta sit
amet tincidunt nec, imperdiet nec lorem.

Mauris lacus lectus, aliquam nec tempus ut, scelerisque placerat erat. Proin
vitae purus pulvinar, feugiat ligula ultrices, ullamcorper tortor. Curabitur
lorem est, vestibulum vel elementum eget, tempor eu massa. Nulla pharetra odio
gravida, luctus magna id, congue ligula. Curabitur porta est eget imperdiet
porta. Fusce at sodales lectus. In lorem lectus, pellentesque at sodales ac,
consectetur nec tellus.

Cras hendrerit rhoncus tellus in venenatis. Quisque auctor est non libero
scelerisque accumsan. Integer molestie, nibh quis pulvinar fermentum, urna erat
consequat ex, at suscipit purus mauris at mi. Etiam posuere elementum dui a
vulputate. Donec risus leo, laoreet ut imperdiet at, euismod nec diam. Vivamus
molestie leo vel nisl consectetur euismod. Fusce mattis porta turpis. Praesent
molestie dolor a lacus scelerisque, eget convallis risus pellentesque. Donec
auctor risus sed quam luctus iaculis. In hac habitasse platea dictumst.
Phasellus venenatis risus non arcu sodales, eu bibendum massa blandit. Quisque
scelerisque augue tortor, eu blandit tortor finibus non.

Donec id ullamcorper leo. In eu tellus justo. Curabitur vel risus malesuada,
mattis nibh vel, tempus purus. Mauris ex mi, rhoncus at ullamcorper vitae,
placerat quis libero. Nullam et sapien leo. Donec efficitur leo sit amet justo
molestie, vel condimentum neque viverra. Nulla quam tortor, posuere iaculis
hendrerit at, luctus a nibh. Suspendisse sagittis in justo quis feugiat. Etiam
accumsan vel mauris eget molestie. Nullam sagittis nulla quis malesuada
fringilla. Maecenas at neque tempor, tincidunt nisi ut, pharetra nibh. Maecenas
id accumsan purus. Nunc porttitor id sapien et efficitur. Morbi fermentum dui
non massa pretium bibendum. Vestibulum mauris sapien, malesuada vitae rhoncus
vitae, venenatis quis nulla. Phasellus maximus, turpis at cursus molestie, elit
lectus mollis eros, nec pharetra nisl magna eget ante.

Aliquam tortor ante, facilisis et lectus aliquet, tempor tincidunt mi. Nunc nec
bibendum neque, sit amet porttitor mauris. Praesent egestas pharetra blandit.
Vestibulum eu dui massa. Vestibulum tincidunt, arcu eget dapibus scelerisque,
nulla eros pretium purus, eu condimentum ligula nibh in mi. Aenean non lorem
eget massa volutpat rutrum id quis libero. Sed vulputate, metus sed suscipit
facilisis, lacus diam luctus mi, vel rhoncus diam magna fermentum urna. Sed
semper egestas diam vel scelerisque. Donec maximus nibh aliquet enim placerat,
eu tincidunt lorem imperdiet. Duis nec odio egestas, hendrerit lectus vel,
pharetra orci. Orci varius natoque penatibus et magnis dis parturient montes,
nascetur ridiculus mus. Suspendisse suscipit turpis sed pretium semper.

Fusce ac dolor porttitor, interdum nulla nec, vestibulum nibh. Vestibulum
gravida tempus velit, id tincidunt nulla accumsan vitae. Etiam gravida orci ac
hendrerit tincidunt. Ut congue enim in nisl sodales, vel aliquam nunc feugiat.
Nulla massa urna, volutpat vel nisi a, placerat commodo augue. Aenean fermentum
sollicitudin elit. Donec vitae purus eget nibh vehicula commodo ac quis odio.
Etiam lacus ante, aliquam in pellentesque at, laoreet non velit. Suspendisse
lacinia libero eu imperdiet luctus. Vestibulum ante ipsum primis in faucibus
orci luctus et ultrices posuere cubilia Curae; Morbi hendrerit facilisis eros ac
posuere. Pellentesque habitant morbi tristique senectus et netus et malesuada
fames ac turpis egestas. Phasellus sagittis justo eget dui condimentum, quis
suscipit magna dictum.

Nullam blandit, mi vel sollicitudin condimentum, nunc odio fringilla tellus, at
aliquet nibh neque quis ligula. Lorem ipsum dolor sit amet, consectetur
adipiscing elit. Donec accumsan dolor ut nibh suscipit bibendum. Vivamus erat
ex, tempus eu ullamcorper sed, facilisis eget arcu. Morbi vel eros vel ex rutrum
vehicula. Ut non massa nisi. Mauris tincidunt pellentesque massa id posuere.
Nullam id mollis turpis. Fusce velit lacus, pretium luctus congue tincidunt,
posuere id dolor. Duis pellentesque purus non diam congue dapibus. Nam ac est
elit. Etiam viverra a diam vitae venenatis. Nullam efficitur tellus non
ultricies mattis. Vestibulum ante ipsum primis in faucibus orci luctus et
ultrices posuere cubilia Curae; Nam in leo eget lacus blandit ullamcorper. Ut
vel lacinia velit, ac cursus urna.

Aenean auctor turpis at tortor tincidunt, id suscipit velit tincidunt. Aliquam
mattis lacus eros, quis porttitor urna suscipit eget. Cras mi tortor, laoreet in
lorem et, porttitor lacinia libero. Nullam aliquet nibh nec sapien auctor, at
pharetra quam viverra. Suspendisse luctus sem arcu, id interdum felis fringilla
quis. In vehicula lectus at purus aliquet, et viverra lacus lobortis. Cras
ultrices, ex ut rhoncus scelerisque, ex risus viverra elit, mattis cursus purus
mi eu ante. Integer at consectetur ligula, non viverra nunc. Duis non vestibulum
tellus. Interdum et malesuada fames ac ante ipsum primis in faucibus.

Aliquam erat volutpat. Mauris sodales pretium turpis, eget tempus sapien
sollicitudin sed. Suspendisse et viverra massa. Morbi arcu neque, feugiat ac
mauris eget, feugiat dapibus eros. Integer blandit auctor tellus, ac consectetur
metus molestie quis. Ut varius leo eget.
`
