import fs from 'fs'

type ConfigMap = Map<string, string>;

async function loadGitignoreConfig (filePath = ".gitignore") {
    const config : ConfigMap = new Map()

    try {
        const fileContent = await fs.promises.readFile(filePath,'utf-8')
        const lines = fileContent.split('\n')

        for (const line of lines) {
           const trimmedline = line.trim()

           if (!trimmedline || trimmedline.startsWith("#")) continue;

            const [key, value] = trimmedline.split("=")
            if(key && value) config.set(key.trim(),value.trim())

        }
       return new Proxy(config,{
        get(target,prop){
            if(typeof prop === 'string' && target.has(prop)){
                return target.get(prop)
            }
            return undefined;
        }
       })
    } catch (error) {
        console.error('Error reading the file:', error);
        return {};
    }
}

const cfg : any = await loadGitignoreConfig()

console.log(cfg.API_KEY)
console.log(typeof cfg.N)
console.log(cfg.NALLA)