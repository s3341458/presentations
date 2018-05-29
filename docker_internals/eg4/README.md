sudo mount -t overlay overlay -o lowerdir=./image_layer1:./image_layer2,upperdir=./container_layer,workdir=./work ./merged

make a merged folder, check image_layer folders and container_layer folder, modify some files check them again

when you want to go further have volume involved:
you can either use overlay file system way to implement docker volume:
sudo mount -t overlay overlay -o lowerdir=./volume,upperdir=./volume,workdir=./work/volume ./merged/volume
or
use mount bind to implement docker volume
sudo mount --bind volume ./merged/volume


unset volume by
sudo umount ./merged/volume

unset merged folder by
sudo umount ./merged/volume

