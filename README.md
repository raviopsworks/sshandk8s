# ssh-cookbook

TODO: Enter the cookbook description here.

berks install
berks upload


knife data bag create user_data_bag
knife data_bag from file user_data_bag user_data_bag/sshuser.json

knife cookbook upload ssh-cookbook
