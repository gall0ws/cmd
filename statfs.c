#include <sys/param.h>
#include <sys/mount.h>
#include <pwd.h>

#include <err.h>
#include <sysexits.h>
#include <stdio.h>

struct mntflags {
	int flag;
	char *name;
};

struct mntflags mntflags[] = {
	{ MNT_LOCAL, "local" },
	{ MNT_RDONLY, "read-only" },
	{ MNT_AUTOMOUNTED, "automounted" },
	{ MNT_JOURNALED, "journaled" },
	{ MNT_NOATIME, "noatime" },
	{ MNT_DONTBROWSE, "nobrowse" },
	{ MNT_NODEV, "nodev" },
	{ MNT_NOEXEC, "noexec" },
	{ MNT_NOSUID, "nosuid" },
	{ MNT_UNION, "union" },
	{ MNT_ASYNC, "async" },
	{ MNT_SYNCHRONOUS, "synchronous" },
	{ MNT_ROOTFS, "rootfs" },
	{ MNT_CPROTECT, "protect" },
	{ MNT_EXPORTED, "exported" },
	{ MNT_REMOVABLE, "removable" },
	{ MNT_QUARANTINE, "quarantine" },
	{ MNT_QUOTA, "quota" },
	{ MNT_IGNORE_OWNERSHIP, "ignore_ownership" },
	{ MNT_NOUSERXATTR, "nouserxattr" },
	{ MNT_DEFWRITE, "defwrite" },
	{ MNT_NOFOLLOW, "nofollow" },
	{ MNT_SNAPSHOT, "snapshot" },
	{ MNT_STRICTATIME, "strictatime" },
	{ 0, NULL }
};

int
main(int argc, char **argv)
{
	struct statfs st;
	struct passwd *pw;
	struct mntflags *p;
	char *path;

	if (argc < 2) {
		fprintf(stderr, "usage: statfs path\n");
		return EX_USAGE;
	}
	path = argv[1];
	if (statfs(path, &st) < 0) {
		err(EX_UNAVAILABLE, "%s", path);
	}

	printf("name: %s\n", st.f_fstypename);
	printf("type: %u.%ud\n", st.f_type, st.f_fssubtype);
	printf("mounted from: %s\n", st.f_mntfromname);
	printf("mounted to: %s\n", st.f_mntonname);

	printf("mount flags: ");
	for (p=mntflags; p->name; p++) {
		if (st.f_flags & p->flag) {
			printf("%s, ", p->name);
		}
	}
	printf("\b\b \n");

	printf("owner: ");
	pw = getpwuid(st.f_owner);
	if (pw) {
		printf("%s (%d)\n", pw->pw_name, st.f_owner);
	} else {
		printf("%d\n", st.f_owner);
	}

	printf("block size: %u\n", st.f_bsize);
	printf("i/o block size: %d\n", st.f_iosize);
	printf("blocks: %llu\n", st.f_blocks);
	printf("blocks free: %llu (%.0f%%)\n", st.f_bfree, (st.f_bfree/(float)st.f_blocks)*100);
	printf("blocks available (non-root): %llu (%.0f%%)\n", st.f_bavail, (st.f_bavail/(float)st.f_blocks)*100);
	printf("files: %llu\n", st.f_files);
	printf("free nodes: %llu\n", st.f_ffree);

	return EX_OK;
}
