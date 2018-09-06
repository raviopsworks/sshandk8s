# Add ssh key to user


users_manage 'sshgroup' do
    group_id 4000
    action [:create]
    data_bag 'user_data_bag'
end