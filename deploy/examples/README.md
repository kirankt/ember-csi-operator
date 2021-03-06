This directory contains a few example driver-specific custom resource files, PVC files and a dummy application YAML file.
- `driver-template.yaml` lists all the specs of the EmberCSI CustomResource and can be used as a starting point to creating driver-specific CustomResource definition.
- `pvc.yaml` is an example PVC file which utilizes the driver-specific StorageClass deployed by Ember-CSI-Operator. 
- `pvc-block.yaml` is an example block volume PVC.
- `snapshot.yaml` is an example of creating snapshots using Ember-CSI
- `restore.yaml` is an example to restore volumes from snapshots.
- `app.yaml` is a dummy application that uses the above PVC.
- A few example backends YAML files are included in the `drivers` directory:
  - `lvm.yaml` for LVM with LIO
  - `ceph.yaml` for RBD/Ceph
  - `ceph-with-topology.yaml` for RBD/Ceph with support for volume-aware topology
  - `xtremio-*.yaml` for Dell EMC XtremIO
  - `vmax.yaml` for Dell EMC VMAX
  - `kaminario.yaml` for Kaminario
  - `solidfire.yaml` for NetApp SolidFire
  - `3par.yaml` for HPE 3PAR
  - `synology.yaml` for Synology
  - `qnap.yaml` for QNAP
