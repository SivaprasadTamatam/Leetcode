#include <algorithm>
#include <vector>

void join(int u, int v, std::vector<int>& parent, std::vector<int>& rank) {
    int p1 = find(u, parent);
    int p2 = find(v, parent);

    if (p1 == p2) {
        return;
    }

    if (rank[p1] > rank[p2]) {
        parent[p2] = p1;
    } else if (rank[p2] > rank[p1]) {
        parent[p1] = p2;
    } else {
        parent[p2] = p1;
        rank[p1]++;
    }
}

int find(int u, std::vector<int>& parent) {
    if (u != parent[u]) {
        parent[u] = find(parent[u], parent);
    }
    return parent[u];
}

bool isConnected(int u, int v, std::vector<int>& parent) {
    return find(parent[u], parent) == find(parent[v], parent);
}

std::vector<bool> distanceLimitedPathsExist(int n, std::vector<std::vector<int>>& edgeList, std::vector<std::vector<int>>& queries) {
    // Add index, so we can fill answer later proper after sorting
    for (int i = 0; i < queries.size(); i++) {
        queries[i].push_back(i);
    }

    // Sort according to distance
    std::sort(edgeList.begin(), edgeList.end(), [](const std::vector<int>& a, const std::vector<int>& b) {
        return a[2] < b[2];
    });

    // Sort according to distance
    std::sort(queries.begin(), queries.end(), [](const std::vector<int>& a, const std::vector<int>& b) {
        return a[2] < b[2];
    });

    std::vector<int> parent(n);
    std::vector<int> rank(n);
    std::vector<bool> res(queries.size());
    for (int i = 0; i < n; i++) {
        parent[i] = i;
    }

    for (int i = 0, j = 0; i < queries.size(); i++) {
        // Join all edges with distance less than d
        while (j < edgeList.size() && edgeList[j][2] < queries[i][2]) {
            join(edgeList[j][0], edgeList[j][1], parent, rank);
            j++;
        }

        // Now check whether vertex in query are connected
        // If distance of all connected edges were lesser, it would have connected in earlier loop
        if (isConnected(queries[i][0], queries[i][1], parent)) {
            res[queries[i][3]] = true;
        } else {
            res[queries[i][3]] = false;
        }
    }

    return res;
}
