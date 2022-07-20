import skimage
import h5py
import os
filepath = r'C:\Users\Administrator\Desktop\ftp\codes\StableNet\StableNet\DATA\mridataset_h5\beijing_gender_64_76_122.hdf5'
f1 = h5py.File(filepath, "r")
img_3d = f1['x'][()] 
f1.close()
verts, faces, normals, values = skimage.measure.marching_cubes_lewiner(img_3d[0,:], 0)

faces=faces +1
thefile = open('brain1.obj', 'w')
for item in verts:
  thefile.write("v {0} {1} {2}\n".format(item[0],item[1],item[2]))

for item in normals:
  thefile.write("vn {0} {1} {2}\n".format(item[0],item[1],item[2]))

for item in faces:
  thefile.write("f {0}//{0} {1}//{1} {2}//{2}\n".format(item[0],item[1],item[2]))  

thefile.close()
os.pause()