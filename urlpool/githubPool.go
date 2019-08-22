package urlpool

var githubReposPool = []bgetFilesURLType{
	{
		Name:         "github_demo",
		Site:         "github",
		URL:          []string{"https://github.com/Miachol/github_demo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bwa",
		Site:         "github",
		URL:          []string{"https://github.com/lh3/bwa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "star",
		Site:         "github",
		URL:          []string{"https://github.com/alexdobin/STAR"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "samtools",
		Site:         "github",
		URL:          []string{"https://github.com/samtools/samtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bcftools",
		Site:         "github",
		URL:          []string{"https://github.com/samtools/bcftools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bowtie",
		Site:         "github",
		URL:          []string{"https://github.com/BenLangmead/bowtie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bowtie2",
		Site:         "github",
		URL:          []string{"https://github.com/BenLangmead/bowtie2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "tophat",
		Site:         "github",
		URL:          []string{"https://github.com/infphilo/tophat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "varscan",
		Site:         "github",
		URL:          []string{"https://github.com/Miachol/varscan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "picard",
		Site:         "github",
		URL:          []string{"https://github.com/broadinstitute/picard"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "vcftools",
		Site:         "github",
		URL:          []string{"https://github.com/vcftools/vcftools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "pindel",
		Site:         "github",
		URL:          []string{"https://github.com/genome/pindel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "lofreq",
		Site:         "github",
		URL:          []string{"https://github.com/Miachol/lofreq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "hisat2",
		Site:         "github",
		URL:          []string{"https://github.com/infphilo/hisat2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "freebayes",
		Site:         "github",
		URL:          []string{"https://github.com/ekg/freebayes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "macs",
		Site:         "github",
		URL:          []string{"https://github.com/taoliu/MACS/"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bedtools2",
		Site:         "github",
		URL:          []string{"https://github.com/arq5x/bedtools2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "sparsehash",
		Site:         "github",
		URL:          []string{"https://github.com/sparsehash/sparsehash"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "abyss",
		Site:         "github",
		URL:          []string{"https://github.com/bcgsc/abyss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bamtools",
		Site:         "github",
		URL:          []string{"https://github.com/pezmaster31/bamtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "breakdancer",
		Site:         "github",
		URL:          []string{"https://github.com/genome/breakdancer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "htseq",
		Site:         "github",
		URL:          []string{"https://github.com/simon-anders/htseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seqtk",
		Site:         "github",
		URL:          []string{"https://github.com/ndaniel/seqtk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "delly",
		Site:         "github",
		URL:          []string{"https://github.com/dellytools/delly"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "tmap",
		Site:         "github",
		URL:          []string{"git://github.com/iontorrent/TMAP"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "somatic-sniper",
		Site:         "github",
		URL:          []string{"https://github.com/genome/somatic-sniper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bamutil",
		Site:         "github",
		URL:          []string{"https://github.com/statgen/bamUtil"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "vcflib",
		Site:         "github",
		URL:          []string{"https://github.com/vcflib/vcflib"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "jvarkit",
		Site:         "github",
		URL:          []string{"https://github.com/lindenb/jvarkit/"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "fastx_toolkit",
		Site:         "github",
		URL:          []string{"https://github.com/agordon/fastx_toolkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "libgtextutils",
		Site:         "github",
		URL:          []string{"https://github.com/agordon/libgtextutils"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "trinityrnaseq",
		Site:         "github",
		URL:          []string{"https://github.com/trinityrnaseq/trinityrnaseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "oases",
		Site:         "github",
		URL:          []string{"https://github.com/dzerbino/oases"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rum",
		Site:         "github",
		URL:          []string{"https://github.com/itmat/rum"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "igraph",
		Site:         "github",
		URL:          []string{"https://github.com/igraph/igraph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "pxz",
		Site:         "github",
		URL:          []string{"https://github.com/jnovy/pxz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "cnvkit",
		Site:         "github",
		URL:          []string{"https://github.com/etal/cnvkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "speedseq",
		Site:         "github",
		URL:          []string{"https://github.com/hall-lab/speedseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "cnvnator",
		Site:         "github",
		URL:          []string{"https://github.com/abyzovlab/CNVnator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "asap",
		Site:         "github",
		URL:          []string{"https://github.com/DeplanckeLab/ASAP"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mdseq",
		Site:         "github",
		URL:          []string{"https://github.com/zjdaye/MDSeq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "sclvm",
		Site:         "github",
		URL:          []string{"https://github.com/PMBio/scLVM"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "f-sclvm",
		Site:         "github",
		URL:          []string{"https://github.com/PMBio/f-scLVM"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bearscc",
		Site:         "github",
		URL:          []string{"https://github.com/Miachol/bearscc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "singlesplice",
		Site:         "github",
		URL:          []string{"https://github.com/jw156605/SingleSplice"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "isop",
		Site:         "github",
		URL:          []string{"https://github.com/nghiavtr/ISOP"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "brie",
		Site:         "github",
		URL:          []string{"https://github.com/huangyh09/brie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "outrigger",
		Site:         "github",
		URL:          []string{"https://github.com/YeoLab/outrigger"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "raceid",
		Site:         "github",
		URL:          []string{"https://github.com/dgrun/RaceID"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "backspin",
		Site:         "github",
		URL:          []string{"https://github.com/linnarsson-lab/BackSPIN"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "zifa",
		Site:         "github",
		URL:          []string{"https://github.com/epierson9/ZIFA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seurat",
		Site:         "github",
		URL:          []string{"https://github.com/satijalab/seurat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rca",
		Site:         "github",
		URL:          []string{"https://github.com/GIS-SP-Group/RCA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mimosca",
		Site:         "github",
		URL:          []string{"https://github.com/asncd/MIMOSCA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "tracer",
		Site:         "github",
		URL:          []string{"https://github.com/teichlab/tracer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "scnorm",
		Site:         "github",
		URL:          []string{"https://github.com/rhondabacher/SCnorm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "sleuth",
		Site:         "github",
		URL:          []string{"https://github.com/pachterlab/sleuth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "oncotator",
		Site:         "github",
		URL:          []string{"https://github.com/broadinstitute/oncotator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "ensembl-vep",
		Site:         "github",
		URL:          []string{"https://github.com/Ensembl/ensembl-vep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "fastq-tools",
		Site:         "github",
		URL:          []string{"https://github.com/dcjones/fastq-tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "arnapipe",
		Site:         "github",
		URL:          []string{"https://github.com/HudsonAlpha/aRNAPipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "trimgalore",
		Site:         "github",
		URL:          []string{"https://github.com/FelixKrueger/TrimGalore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "multiqc",
		Site:         "github",
		URL:          []string{"https://github.com/ewels/MultiQC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "strawberry",
		Site:         "github",
		URL:          []string{"https://github.com/ruolin/strawberry"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "fastp",
		Site:         "github",
		URL:          []string{"https://github.com/OpenGene/fastp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "olego",
		Site:         "github",
		URL:          []string{"https://github.com/chaolinzhanglab/olego"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "chronqc",
		Site:         "github",
		URL:          []string{"https://github.com/nilesh-tawari/ChronQC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "dart",
		Site:         "github",
		URL:          []string{"https://github.com/hsinnan75/DART"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rhat",
		Site:         "github",
		URL:          []string{"https://github.com/HIT-Bioinformatics/rHAT"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "giggle",
		Site:         "github",
		URL:          []string{"https://github.com/ryanlayer/giggle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "facets",
		Site:         "github",
		URL:          []string{"https://github.com/mskcc/facets"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "resm",
		Site:         "github",
		URL:          []string{"https://github.com/deweylab/RSEM"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "radia",
		Site:         "github",
		URL:          []string{"https://github.com/aradenbaugh/radia/"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "manta",
		Site:         "github",
		URL:          []string{"https://github.com/Illumina/manta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "jaffa",
		Site:         "github",
		URL:          []string{"https://github.com/Oshlack/JAFFA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "marvel",
		Site:         "github",
		URL:          []string{"https://github.com/schloi/MARVEL"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "picky",
		Site:         "github",
		URL:          []string{"https://github.com/TheJacksonLaboratory/Picky"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "freec",
		Site:         "github",
		URL:          []string{"https://github.com/BoevaLab/FREEC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "taxmaps",
		Site:         "github",
		URL:          []string{"https://github.com/nygenome/taxmaps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "svaba",
		Site:         "github",
		URL:          []string{"https://github.com/walaj/svaba"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rop",
		Site:         "github",
		URL:          []string{"https://github.com/smangul1/rop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mrnn",
		Site:         "github",
		URL:          []string{"https://github.com/hendrixlab/mRNN"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "dna-seq-gatk-variant-calling",
		Site:         "github",
		URL:          []string{"https://github.com/snakemake-workflows/dna-seq-gatk-variant-calling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "spack",
		Site:         "github",
		URL:          []string{"https://github.com/spack/spack"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bpipe",
		Site:         "github",
		URL:          []string{"https://github.com/ssadedin/bpipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mscentipede",
		Site:         "github",
		URL:          []string{"https://github.com/rajanil/msCentipede"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "chia-pet2",
		Site:         "github",
		URL:          []string{"https://github.com/GuipengLi/ChIA-PET2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "chromtime",
		Site:         "github",
		URL:          []string{"https://github.com/ernstlab/ChromTime"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "dreg",
		Site:         "github",
		URL:          []string{"https://github.com/Danko-Lab/dREG"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "genomedisco",
		Site:         "github",
		URL:          []string{"https://github.com/kundajelab/genomedisco"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
}
