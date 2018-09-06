# ssh-cookbook
  
ssh-cookbook to add ssh keys for a user in dev environment using roles and data_bags<br/>
<br/>
berks install<br/>
berks upload<br/>
<br/>
<br/>

knife data bag create user_data_bag<br/>
knife data_bag from file user_data_bag user_data_bag/sshuser.json<br/>

<br/>

knife cookbook upload ssh-cookbook<br/>