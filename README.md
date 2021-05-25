# shipa-terraform-provider


# Building and testing

1. Build terraform provider by using:
```
    make install
```

2. Set env values:
```    
    export SHIPA_HOST=http://target.shipa.cloud:8080
    export SHIPA_TOKEN=xxxxxxxxx
```

3. Run terraform
```
    cd example
    terraform init && terraform apply --auto-approve  
``` 
