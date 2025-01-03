# TIG

This CLI tool simplifies managing Git configurations across multiple repositories. It is organized by 'Organizations,' each containing all the essential information, making it easy to switch between SSH credentials and repository settings. Ideal for developers working with different credentials in various projects, this tool ensures seamless transitions and efficient workflow management.

## Steps for installation

**0. Install dependencies**
```
sudo apt install openssh-client
sudo apt install git
```

**1. Download the executable file**
```
wget https://github.com/sneuder/tig/releases/download/v2.0.0/install.sh
```

**2. Make the Script Executable**
```
sudo chmod +x install.sh
```

**3. Run the Script**
```
sudo ./install.sh
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
tig checkout <organization_name>
```

- Remove tig organization
```
tig remove <organization_name>
```

- Rebuild tig organization's credentials
```
tig ssh <organization_name>
```

- Export tig organization's in json file
```
tig export
```

- Build tig organization from json file
```
tig build
```

- Help
```
tig help
```
