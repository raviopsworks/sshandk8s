name 'ssh-cookbook'
maintainer 'Ravi Teja'
maintainer_email 'raviopsworks@gmail.com'
license 'All Rights Reserved'
description 'Installs/Configures ssh-cookbook'
long_description 'Install ssh key to all servers for a particular user'
version '0.1.0'
chef_version '>= 12.14' if respond_to?(:chef_version)

depends 'users'
depends 'sudo'
