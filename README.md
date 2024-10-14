# TIG

This CLI tool simplifies managing Git configurations across multiple repositories. It is organized by 'Organizations,' each containing all the essential information, making it easy to switch between SSH credentials and repository settings. Ideal for developers working with different credentials in various projects, this tool ensures seamless transitions and efficient workflow management.

## Steps for installation

**1. Download the executable file**
```
wget https://github.com/tig.tar.gz
```


**2. Unpack the downloaded file**
```
tar -xzf tig.tar.gz
```

## Commands

- Add tig organization
```
tig add -o tig-sample -e sneuder@tig.com -n sneuder -p github
```

- List tig organization
```
tig ls
```

- Checkout tig organization
```
tig checkout <organization name>
```

- Remove tig organization
```
tig remove <organization name>
```

- Rebuild tig organization's credentials
```
tig ssh <organization name>
```

- Help
```
tig help
```