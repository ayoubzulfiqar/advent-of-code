import networkx as nx


def getIntoLANParty() -> str:
    g = nx.Graph()

    with open("input.txt") as f:
        for lines in f.read().splitlines():
            ps = lines.split("-")
            g.add_edge(ps[0], ps[1])

    clusters = list(nx.find_cliques_recursive(g))
    clusters.sort(key=len, reverse=True)
    clusters[0].sort()
    cluster: str = ",".join(clusters[0])
    print(cluster)
    return cluster


# ai,bk,dc,dx,fo,gx,hk,kd,os,uz,xn,yk,zs
